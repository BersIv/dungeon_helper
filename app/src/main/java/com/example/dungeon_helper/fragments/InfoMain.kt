package com.example.dungeon_helper.fragments

import androidx.lifecycle.ViewModelProvider
import android.os.Bundle
import androidx.fragment.app.Fragment
import android.view.LayoutInflater
import android.view.View
import android.view.ViewGroup
import android.widget.TextView
import com.example.dungeon_helper.R
import com.example.dungeon_helper.databinding.FragmentInfoMainBinding

class InfoMain : Fragment() {

    companion object {
        fun newInstance() = InfoMain()
    }
    private lateinit var viewModel: InfoMainViewModel


    private var _binding: FragmentInfoMainBinding? = null
    private  val binding get() = _binding!!

    override fun onCreateView(
        inflater: LayoutInflater,
        container: ViewGroup?,
        savedInstanceState: Bundle?
    ): View {
        val infoMainViewModel = ViewModelProvider(this)[InfoMainViewModel::class.java]

        _binding = FragmentInfoMainBinding.inflate(inflater, container, false)
        val root: View = binding.root

        val textView: TextView = binding.textInfo
        infoMainViewModel.text.observe(viewLifecycleOwner) {
            textView.text = it
        }

        return root
    }

    override fun onDestroyView() {
        super.onDestroyView()
        _binding = null
    }

}