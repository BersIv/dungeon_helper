package com.example.dungeon_helper.fragments.main

import androidx.lifecycle.ViewModelProvider
import android.os.Bundle
import androidx.fragment.app.Fragment
import android.view.LayoutInflater
import android.view.View
import android.view.ViewGroup
import android.widget.TextView
import com.example.dungeon_helper.MainActivity
import com.example.dungeon_helper.R
import com.example.dungeon_helper.databinding.FragmentInfoVersBinding

class InfoVers : Fragment() {

    companion object {
        fun newInstance() = InfoVers()
    }

    private lateinit var viewModel: InfoVersViewModel

    private var _binding: FragmentInfoVersBinding? = null
    private val binding get() = _binding!!
    private lateinit var gameVersionView: TextView
    private lateinit var gameTitleView: TextView

    override fun onCreateView(
        inflater: LayoutInflater,
        container: ViewGroup?,
        savedInstanceState: Bundle?
    ): View? {
       val infoVersViewModel = ViewModelProvider(this)[InfoVersViewModel::class.java]
        _binding = FragmentInfoVersBinding.inflate(inflater, container,false)
        val root: View = binding.root
        val textView:TextView = binding.textInfoVers
        infoVersViewModel.text.observe(viewLifecycleOwner){
            textView.text = it
        }
        return root
    }

    override fun onDestroyView() {
        super.onDestroyView()
        _binding = null
    }

    override fun onStart() {
        super.onStart()
        gameTitleView = binding.gameTitle
        gameTitleView.setOnClickListener {
            (activity as MainActivity).navController.navigate(R.id.action_infoVers_to_navigation_info)
        }
        gameVersionView = binding.gameVersion
        gameVersionView.setOnClickListener {
            (activity as MainActivity).navController.navigate(R.id.action_infoVers_to_infoAllChapters)
        }

    }

}