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
import com.example.dungeon_helper.databinding.FragmentInfoAllChaptersBinding

class InfoAllChapters : Fragment() {

    companion object {
        fun newInstance() = InfoAllChapters()
    }

    private lateinit var viewModel: InfoAllChaptersViewModel
    private var _binding: FragmentInfoAllChaptersBinding? = null
    private val binding get() = _binding!!
    private lateinit var gameVersionView: TextView
    private lateinit var gameTitleView: TextView
    private lateinit var gameChaptersView: TextView
    override fun onCreateView(
        inflater: LayoutInflater, container: ViewGroup?,
        savedInstanceState: Bundle?
    ): View? {
        val infoAllChaptersViewModel = ViewModelProvider(this)[InfoAllChaptersViewModel::class.java]
        _binding = FragmentInfoAllChaptersBinding.inflate(inflater, container,false)
        val root: View = binding.root
        val textView:TextView = binding.textInfoChapter
        infoAllChaptersViewModel.text.observe(viewLifecycleOwner){
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
            (activity as MainActivity).navController.navigate(R.id.action_infoAllChapters_to_navigation_info)
        }
        gameVersionView = binding.gameVersion
        gameVersionView.setOnClickListener {
            (activity as MainActivity).navController.navigate(R.id.action_infoAllChapters_to_infoVers)
        }
        gameChaptersView = binding.gameChapter1
        gameChaptersView.setOnClickListener {
            (activity as MainActivity).navController.navigate(R.id.action_infoAllChapters_to_infoChapter1Content)
        }

    }


}