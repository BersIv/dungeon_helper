package com.example.dungeon_helper.fragments.lobby

import androidx.lifecycle.ViewModelProvider
import android.os.Bundle
import androidx.fragment.app.Fragment
import android.view.LayoutInflater
import android.view.View
import android.view.ViewGroup
import android.widget.TextView
import com.example.dungeon_helper.R
import com.example.dungeon_helper.databinding.FragmentLobbyBinding

class Lobby : Fragment() {

    companion object {
        fun newInstance() = Lobby()
    }

    private lateinit var viewModel: LobbyViewModel

    private var _binding: FragmentLobbyBinding? = null
    private  val binding get() = _binding!!

    override fun onCreateView(
        inflater: LayoutInflater,
        container: ViewGroup?,
        savedInstanceState: Bundle?
    ): View? {
        val infoMainViewModel = ViewModelProvider(this)[LobbyViewModel::class.java]

        _binding = FragmentLobbyBinding.inflate(inflater, container, false)
        val root: View = binding.root

        val textView: TextView = binding.textPlay
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